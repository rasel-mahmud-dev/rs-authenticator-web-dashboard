import {createBrowserRouter} from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute';
import NotFound from "../components/NotFound.jsx";
import LoginForm from "../components/LoginForm.jsx";
import HomePage from "../pages/HomePage.jsx";
import Layout from "../components/Layout.jsx";
import ConnectedAuthenticators from "../pages/ConnectedAuthenticators.jsx";
import AuthenticatorSetup from "../pages/AuthenticatorSetup.jsx";
import SetupGoogleAuthenticator from "../pages/SetupGoogle.jsx";
import AccountLayout from "../components/AccountLayout.jsx";
import TwoFactorAuthentication from "../pages/TwoFactorAuthentication.jsx";

const routes = createBrowserRouter([
    {
        path: '/',
        element: <Layout/>,
        errorElement: <NotFound/>,
        children: [
            {
                path: 'login',
                element: <LoginForm/>,
            },
            {
                path: 'login/two-factor-authenticator',
                element: <TwoFactorAuthentication/>,
            },
            {

                path: "/account",
                element: <AccountLayout/>,
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
                        path: "authenticator-apps",
                        element: <ProtectedRoute><ConnectedAuthenticators/></ProtectedRoute>
                    }, {
                        path: "authenticator-setup",
                        element: <ProtectedRoute><AuthenticatorSetup/></ProtectedRoute>
                    },
                    {
                        path: "setup-google-auth",
                        element: <ProtectedRoute><SetupGoogleAuthenticator/></ProtectedRoute>
                    }
                ]
            }

            // <Route path="/setup-microsoft-auth" element={ <ProtectedRoute><MicrosoftAuthenticatorSetupPage />} />
        ],
    },
]);

export default routes;
