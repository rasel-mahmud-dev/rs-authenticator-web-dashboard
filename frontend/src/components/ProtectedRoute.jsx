import { Navigate } from 'react-router-dom';
import useAuthStore from "../store/authState.js";

const ProtectedRoute = ({ children }) => {
    const {  user, authLoaded } = useAuthStore()

    if (!authLoaded) {
        return <div>Loading...</div>;
    }

    if (!user?.id) {
        return <Navigate to="/login" replace />;
    }

    return <>{children}</>;
};

export default ProtectedRoute;
