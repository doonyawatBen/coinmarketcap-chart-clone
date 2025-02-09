import { FC } from 'react';
import { useRouteError } from "react-router-dom";

interface ErrorPageProps {
}

const ErrorPage: FC<ErrorPageProps> = () => {
    const error: any = useRouteError();
    console.error(error);

    return (
        <div id="error-page" className='flex flex-col items-center self-center'>
            <h1>Oops!</h1>
            <p>Sorry, an unexpected error has occurred.</p>
            <p>
                <i>{error.statusText || error.message}</i>
            </p>
        </div>
    );
}

export default ErrorPage;