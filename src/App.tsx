import {
  RouterProvider,
  createBrowserRouter,
} from "react-router-dom";

import ErrorPage from "./error-page";
import Graph from "./Graph";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Graph/>,
    errorElement:<ErrorPage/>
  },
]);

function App() {
  return (
    <RouterProvider router={router} />
  );
}

export default App;
