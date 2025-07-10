import type { JSX } from "react";
import { Navigate } from "react-router-dom";

const AuthRoute = ({ children }: { children: JSX.Element }) =>{
    const token = localStorage.getItem("token");
    return token ? children : <Navigate to="/login" replace />;
};

export default AuthRoute;