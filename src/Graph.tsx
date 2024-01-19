import {
  AxisLabelsLocation,
  GridMode,
  IgrChartMouseEventArgs,
  IgrCrosshairLayer,
  IgrCrosshairLayerModule,
  IgrDataChart,
  IgrDataChartAnnotationModule,
  IgrDataChartCategoryModule,
  IgrDataChartCoreModule,
  IgrDataChartInteractivityModule,
  IgrNumberAbbreviatorModule,
  IgrNumericYAxis,
  IgrSeriesViewer,
  IgrSplineAreaSeries,
  IgrTimeXAxis,
  IgrZoomSlider,
  IgrZoomSliderModule,
  MarkerType,
} from "igniteui-react-charts";
import { IChartTooltipProps, IgRect, Visibility } from "igniteui-react-core";
import React, { useEffect, useRef, useState } from "react";

import { BITCOIN_DATA } from "./mock-data/bitcoin-data.ts";
import CustomTooltip from "./components/CustomTooltip.tsx";
import { Timeframe } from "./types.ts";
import clsx from "clsx";
import dayjs from "dayjs";
import { formatYAxisValue } from "./utils.ts";
import tailwindConfig from "../tailwind.config.ts";
import { useLocation } from "react-router-dom";
import utc from "dayjs/plugin/utc";

IgrZoomSliderModule.register();
IgrDataChartCoreModule.register();
IgrDataChartCategoryModule.register();
IgrDataChartInteractivityModule.register();
IgrDataChartAnnotationModule.register();
IgrNumberAbbreviatorModule.register();
IgrCrosshairLayerModule.register();

dayjs.extend(utc);
const MINIMAP_HEIGHT = "66px";
const themeColors = tailwindConfig.theme.extend.colors;

function useQuery() {
  const { search } = useLocation();
  return React.useMemo(() => new URLSearchParams(search), [search]);
}

function Graph() {
  const [timeframe, setTimeframe] = useState<Timeframe>("all");
  const mainChartRef = useRef<IgrDataChart | null>(null);
  const zoomSliderRef = useRef<IgrZoomSlider | null>(null);
  const query = useQuery();

  const handleMouseEnter = (
    _s: IgrSeriesViewer,
    chartEvent: IgrChartMouseEventArgs
  ) => {
    chartEvent.series.tooltipTemplate = ({
      dataContext,
    }: IChartTooltipProps) => {
      const { item } = dataContext;
      return <CustomTooltip data={item} />;
    };
  };

  const calculateZoomSliderWindow = (timeframe: Timeframe) => {
    const mainChart = mainChartRef.current;
    const zoomSlider = zoomSliderRef.current;

    if (!mainChart || !zoomSlider) return;

    const allData = BITCOIN_DATA["all"];
    const timeframeData = BITCOIN_DATA[timeframe];

    const startDate = timeframeData[0].date;
    let startPosition = allData.length - 1;

    // we loop backwards to find the earliest date in the allData array that is before the start date of the timeframe
    for (let i = startPosition; i >= 0; i--) {
      startPosition = i;
      if (dayjs(allData[i].date).isBefore(startDate)) {
        break;
      }
    }

    // once we have the start position, we can calculate the width of the zoom slider window
    const width = (allData.length - startPosition) / allData.length;

    zoomSlider.windowRect = {
      ...zoomSlider.windowRect,
      width,
      left: 1 - width,
    };

    const sliderWindow = zoomSlider.windowRect;
    const chartWindow = mainChart.actualWindowRect;

    const zoom = {
      top: chartWindow.top,
      left: sliderWindow.left,
      width: sliderWindow.width,
      height: chartWindow.height,
    };

    setTimeframe(timeframe);
    updateMainChartZoom(zoom);
  };

  // we want to update the main chart zoom when the user changes the zoom slider window
  const handleZoomSliderWindowChanged = () => {
    const mainChart = mainChartRef.current;
    const zoomSlider = zoomSliderRef.current;

    if (!mainChart || !zoomSlider) return;

    // because the zoom charts uses all the historical data, we need to set the main chart source to historical too
    // to make sure the zoom chart is in sync with the main chart
    setTimeframe("all");

    const sliderWindow = zoomSlider.windowRect;
    const chartWindow = mainChart.actualWindowRect;

    const zoom = {
      top: chartWindow.top,
      left: sliderWindow.left,
      width: sliderWindow.width,
      height: chartWindow.height,
    };

    updateMainChartZoom(zoom);
  };

  const updateMainChartZoom = (zoom: IgRect) => {
    const data = BITCOIN_DATA["all"];
    const mainChart = mainChartRef.current;
    const zoomSlider = zoomSliderRef.current;

    if (!mainChart || !zoomSlider) return;

    const yAxis = mainChart.actualAxes[1] as IgrNumericYAxis;

    // we get the indexes using the current left and width of the zoom slider window
    let indexStart = Math.floor((data.length - 1) * zoom.left);
    let indexEnd = Math.ceil((data.length - 1) * (zoom.left + zoom.width));

    let min = Number.MAX_VALUE;
    let max = Number.MIN_VALUE;

    if (indexStart < 0) {
      indexStart = 0;
    }

    indexEnd = Math.min(indexEnd, data.length - 1);

    // once we have the indexes we calculate the min and max prices
    for (let i = indexStart; i <= indexEnd; i++) {
      min = Math.min(min, data[i].price);
      max = Math.max(max, data[i].price);
    }

    // with the price range we can calculate the min and max values for the y-axis
    const yMin =
      (min - yAxis.actualMinimumValue) /
      (yAxis.actualMaximumValue - yAxis.actualMinimumValue);

    const yMax =
      (max - yAxis.actualMinimumValue) /
      (yAxis.actualMaximumValue - yAxis.actualMinimumValue);

    mainChart.windowRect = {
      left: zoom.left,
      width: zoom.width,
      top: 1 - yMax,
      height: yMax - yMin,
    };
  };

  const saveBase64AsHTML = (base64: string) => {
    document.body.style.margin = "0px"
    document.body.style.height = "100%"
    document.body.style.backgroundColor = "rgb(14, 14, 14)"

    const imgCheck = document.getElementById('image-render-graph')
    if (!imgCheck) {
      const img = document.createElement("img");
      document.body.appendChild(img);
      img.id = "image-render-graph"
      img.style.display = "block"
      img.style.margin = "auto"
      img.style.cursor = "zoom-in"
      img.style.backgroundColor = "hsl(0, 0%, 100%)"
      img.style.transition = "background-color 300ms"
      img.src = base64
    }

    const root = document.getElementById('root')
    root?.remove()
    const rootScript = document.getElementById('root-script')
    rootScript?.remove()
  }

  const renderChartToImage = (width: number, height: number) => {
    const image = mainChartRef.current?.renderToImage(width, height);
    saveBase64AsHTML(image.src);
  }

  const calculateZoomSliderWindowParam = (startDate: Date, endDate: Date) => {
    const mainChart = mainChartRef.current;
    const zoomSlider = zoomSliderRef.current;

    if (!mainChart || !zoomSlider) return;

    const allData = BITCOIN_DATA["all"];

    let startPosition = allData.length - 1;

    // Find the index of the start date in the allData array
    for (let i = startPosition; i >= 0; i--) {
      startPosition = i;
      if (dayjs(allData[i].date).isBefore(startDate)) {
        break;
      }
    }

    // Find the index of the end date in the allData array
    let endPosition = allData.length - 1;
    for (let i = endPosition; i >= 0; i--) {
      endPosition = i;
      if (dayjs(allData[i].date).isBefore(endDate)) {
        break;
      }
    }

    // Calculate the width of the zoom slider window based on the start and end positions
    const width = (endPosition - startPosition + 2) / allData.length;
    const left = (1 / (allData.length - 2)) * startPosition

    zoomSlider.windowRect = {
      ...zoomSlider.windowRect,
      width,
      left: left,
    };

    const sliderWindow = zoomSlider.windowRect;
    const chartWindow = mainChart.actualWindowRect;

    const zoom = {
      top: chartWindow.top,
      left: sliderWindow.left,
      width: sliderWindow.width,
      height: chartWindow.height,
    };

    updateMainChartZoom(zoom);
  };

  useEffect(() => {
    const startDateString: string = query.get("startDate") || ""
    const startDate = new Date(startDateString)

    const endDateString: string = query.get("endDate") || ""
    const endDate = new Date(endDateString)

    const widthString = query.get("width") || ""
    const width = parseInt(widthString)

    const heightString = query.get("height") || ""
    const height = parseInt(heightString)

    if (startDateString && endDateString && widthString && heightString) {
      calculateZoomSliderWindowParam(startDate, endDate)

      setTimeout(() => {
        renderChartToImage(width, height)
      }, 0);
    }
  })

  return (
    <div className="w-[929px] h-[396px] mx-auto mt-[100px] px-10">
      <div className="flex flex-row items-center">
        <ul className="flex items center bg-gray-50 p-1 rounded m-0 text-xs text-gray-700 gap-1 ml-auto mb-4">
          <button
            className={clsx(
              "px-[8px] py-[5px] hover:bg-gray-100 hover:font-medium hover:rounded",
              timeframe === "1d" && "bg-white font-medium rounded"
            )}
            onClick={() => calculateZoomSliderWindow("1d")}
          >
            1D
          </button>
          <button
            className={clsx(
              "px-[8px] py-[5px] hover:bg-gray-100 hover:font-medium hover:rounded",
              timeframe === "1w" && "bg-white font-medium rounded"
            )}
            onClick={() => calculateZoomSliderWindow("1w")}
          >
            1W
          </button>
          <button
            className={clsx(
              "px-[8px] py-[5px] hover:bg-gray-100 hover:font-medium hover:rounded",
              timeframe === "1m" && "bg-white font-medium rounded"
            )}
            onClick={() => calculateZoomSliderWindow("1m")}
          >
            1M
          </button>
          <button
            className={clsx(
              "px-[8px] py-[5px] hover:bg-gray-100 hover:font-medium hover:rounded",
              timeframe === "1y" && "bg-white font-medium rounded"
            )}
            onClick={() => calculateZoomSliderWindow("1y")}
          >
            1Y
          </button>
          <button
            className={clsx(
              "px-[8px] py-[5px] hover:bg-gray-100 hover:font-medium hover:rounded",
              timeframe === "all" && "bg-white font-medium rounded"
            )}
            onClick={() => calculateZoomSliderWindow("all")}
          >
            ALL
          </button>
        </ul>
      </div>
      <IgrDataChart
        ref={mainChartRef}
        dataSource={BITCOIN_DATA["all"]}
        width="100%"
        height={`calc(100% - ${MINIMAP_HEIGHT})`}
        defaultInteraction="DragPan"
        seriesMouseEnter={handleMouseEnter}
        outlines={[themeColors.green["500"]]}
        brushes={
          [
            {
              type: "linearGradient",
              colorStops: [
                {
                  color: themeColors.green["50"],
                  offset: 0,
                },
                {
                  color: "white",
                  offset: 0.95,
                },
              ],
            },
          ] as any
        }
      >
        <IgrTimeXAxis
          name="xAxis"
          dateTimeMemberPath="date"
          titleLocation="outsideBottom"
          labelLocation="outsideBottom"
          stroke={themeColors.gray["20"]}
          labelTextColor={themeColors.gray["700"]}
        />
        <IgrNumericYAxis
          name="yAxis"
          abbreviateLargeNumbers={true}
          labelLocation={AxisLabelsLocation.OutsideRight}
          majorStroke={themeColors.gray["10"]}
          stroke={themeColors.gray["20"]}
          labelTextColor={themeColors.gray["700"]}
          formatLabel={(value) => formatYAxisValue(value)}
          title="USD"
          titleTextColor={themeColors.gray["700"]}
          titleLeftMargin={5}
        />
        <IgrCrosshairLayer
          name="CrosshairLayer"
          thickness={1.1}
          verticalLineStroke={themeColors.gray["200"]}
          horizontalLineStroke={themeColors.gray["200"]}
          yAxisAnnotationBackground={themeColors.gray["600"]}
          xAxisAnnotationBackground={themeColors.gray["600"]}
          yAxisAnnotationBackgroundCornerRadius={0}
          xAxisAnnotationBackgroundCornerRadius={0}
          isAxisAnnotationEnabled={true}
        />
        <IgrSplineAreaSeries
          name="priceSeries"
          xAxisName="xAxis"
          yAxisName="yAxis"
          valueMemberPath="price"
          title="Price"
          showDefaultTooltip="false"
          markerType={MarkerType.None}
        />
      </IgrDataChart>
      <div className={`w-full h-[${MINIMAP_HEIGHT}] relative mt-4`}>
        <div className={`w-full h-[${MINIMAP_HEIGHT}] absolute top-0 left-0`}>
          <IgrDataChart
            dataSource={BITCOIN_DATA["all"]}
            width="100%"
            height={MINIMAP_HEIGHT}
            outlines={themeColors.gray["20"]}
            brushes={themeColors.gray["20"]}
            gridMode={GridMode.None}
          >
            <IgrTimeXAxis
              name="zoomXAxis"
              dateTimeMemberPath="date"
              labelTextColor={themeColors.gray["700"]}
            />
            <IgrNumericYAxis
              name="zoomYAxis"
              labelVisibility={Visibility.Collapsed}
            />
            <IgrSplineAreaSeries
              name="priceSeries"
              xAxisName="zoomXAxis"
              yAxisName="zoomYAxis"
              valueMemberPath="price"
              title="Price"
              markerType={MarkerType.None}
            />
          </IgrDataChart>
        </div>
        <div
          // not using class because the canvas of the slider does not inherit the classnames
          style={{
            width: "100%",
            height: MINIMAP_HEIGHT,
            position: "absolute",
            top: "0px",
            left: "0px",
          }}
        >
          <IgrZoomSlider
            ref={zoomSliderRef}
            width="100%"
            height="100%"
            windowRectChanged={handleZoomSliderWindowChanged}
            barExtent={0}
          />
        </div>
      </div>
    </div>
  );
}

export default Graph;
