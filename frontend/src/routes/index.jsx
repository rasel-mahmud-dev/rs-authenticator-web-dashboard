import {createBrowserRouter} from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute';
import NotFound from "../components/NotFound.jsx";
import LoginForm from "../components/LoginForm.jsx";
import HomePage from "../pages/HomePage.jsx";
import Layout from "../components/Layout.jsx";

const routes = createBrowserRouter([
    {
        path: '/',
        element: <Layout/>,
        errorElement: <NotFound/>,
        children: [
            {
                path: '',
                element: (
                    <ProtectedRoute>
                        <HomePage/>
                    </ProtectedRoute>
                ),
            },
            {
                path: 'login',
                element: <LoginForm/>,
            },
        ],
    },
]);

export default routes;
