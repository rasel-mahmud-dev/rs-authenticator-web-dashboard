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
import AuthenticationLogin from "../pages/AuthenticationLogin.jsx";
import BlogHomePage from "../pages/BlogHomePage.jsx";
import ProfilePage from "../pages/ProfilePage.jsx";
import RegistrationForm from "../components/RegistrationForm.jsx";

const routes = createBrowserRouter([
    {
        path: '/',
        element: <Layout/>,
        notFoundElement: <NotFound/>,
        children: [
            {
                path: '',
                element: <BlogHomePage/>,
            },
            {
                path: 'login',
                element: <LoginForm/>,
            },
            {
                path: 'registration',
                element: <RegistrationForm/>,
            },
            {
                path: 'login/authenticator',
                element: <AuthenticationLogin/>,
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
                        path: '/account/profile',
                        element: (
                            <ProtectedRoute>
                                <ProfilePage/>
                            </ProtectedRoute>
                        ),
                    },

                    {
                        path: "/account/authenticator-apps",
                        element: <ProtectedRoute><ConnectedAuthenticators/></ProtectedRoute>
                    },
                    {
                        path: "/account/authenticator-setup",
                        element: <ProtectedRoute><AuthenticatorSetup/></ProtectedRoute>
                    },
                    {
                        path: "/account/authenticator-setup/:provider",
                        element: <ProtectedRoute><SetupGoogleAuthenticator/></ProtectedRoute>
                    },
                ]
            }

            // <Route path="/setup-microsoft-auth" element={ <ProtectedRoute><MicrosoftAuthenticatorSetupPage />} />
        ],
    },
]);

export default routes;
