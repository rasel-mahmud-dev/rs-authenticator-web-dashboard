import {createBrowserRouter} from 'react-router-dom';
import ProtectedRoute from '../components/ProtectedRoute';
import NotFound from "../components/NotFound.jsx";
import LoginForm from "../components/LoginForm.jsx";
import DashboardPage from "../pages/DashboardPage.jsx";
import Layout from "../components/Layout.jsx";
import ConnectedAuthenticators from "../pages/ConnectedAuthenticators.jsx";
import SetupGoogleAuthenticator from "../pages/SetupAuthenticator.jsx";
import AccountLayout from "../components/AccountLayout.jsx";
import AuthenticationLogin from "../pages/AuthenticationLogin.jsx";
import BlogHomePage from "../pages/BlogHomePage.jsx";
import ProfilePage from "../pages/ProfilePage.jsx";
import RegistrationForm from "../components/RegistrationForm.jsx";
import EditProfilePage from "../pages/EditProfilePage.jsx";
import AccountSettings from "../pages/AccountSettings.jsx";
import UsersList from "../pages/UsersList.jsx";
import AboutMe from "../pages/AboutMe.jsx";
import DownloadApp from "../pages/DownloadApp.jsx";
import RecoveryAccount from "../components/RecoveryAccount.jsx";

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
                path: 'account-recovery',
                element: <RecoveryAccount/>,
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
                path: 'about-me',
                element: <AboutMe/>,
            },
            {
                path: 'download-app',
                element: <DownloadApp/>,
            },
            {

                path: "/account",
                element: <AccountLayout/>,
                children: [
                    {
                        path: '',
                        element: (
                            <ProtectedRoute>
                                <DashboardPage/>
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
                    }, {
                        path: '/account/settings',
                        element: (
                            <ProtectedRoute>
                                <AccountSettings/>
                            </ProtectedRoute>
                        ),
                    },
                    {
                        path: '/account/profile/edit',
                        element: (
                            <ProtectedRoute>
                                <EditProfilePage/>
                            </ProtectedRoute>
                        ),
                    },
                    {
                        path: '/account/users',
                        element: (
                            <ProtectedRoute>
                                <UsersList/>
                            </ProtectedRoute>
                        ),
                    },

                    {
                        path: "/account/authenticator-apps",
                        element: <ProtectedRoute><ConnectedAuthenticators/></ProtectedRoute>
                    },

                    {
                        path: "/account/authenticator-setup",
                        element: <ProtectedRoute><SetupGoogleAuthenticator/></ProtectedRoute>
                    },
                ]
            }

            // <Route path="/setup-microsoft-auth" element={ <ProtectedRoute><MicrosoftAuthenticatorSetupPage />} />
        ],
    },
]);

export default routes;
