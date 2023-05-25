import { useLocation, Navigate } from "react-router-dom";

const RequireAuth = (props) => {
    const location = useLocation()

    if (!props.auth) {
        return <Navigate to="/auth/login" state={{from: location}}/>
    }

    if (props.user.group_id !== 1){
        return <Navigate to="/" state={{from: location}}/>
    }

    return props.children
}

export default RequireAuth;