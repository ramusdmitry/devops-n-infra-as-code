import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { getAuth } from './utils';

import { BrowserRouter, Routes, Route } from "react-router-dom"
import FeedPage from "./Pages/Feed/FeedPage";
import DashboardPage from "./Pages/Dashboard/DashboardPage";
import ChatPage from "./Pages/Main";
import LoginPage from './Pages/Auth/LoginPage';
import RegisterPage from './Pages/Auth/RegisterPage';
import Header from './Components/Header';

import NotFoundPage from './Pages/Errors/NotFoundPage';
import LogoutPage from './Pages/Auth/LogoutPage';
import RequireAuth from './Components/Auth/RequireAuth';

function App() {
  const [isLoggedIn, user] = getAuth(); // Request Auth status and user on every render

  return (
      <div>
        <BrowserRouter>
          <Routes>
            <Route
                exact
                path="/"
                element={[<Header auth={isLoggedIn} user={user} />, <ChatPage />]}
            />
            <Route
                exact
                path="/feed"
                element={[<Header auth={isLoggedIn} user={user} />, <FeedPage />]}
            />
            <Route
                exact
                path="/dashboard"
                element={[
                  <Header auth={isLoggedIn} user={user} />,
                  <RequireAuth auth={isLoggedIn} user={user}>
                    <DashboardPage />
                  </RequireAuth>,
                ]}
            />
            <Route exact path="/auth/register" element={<RegisterPage />} />
            <Route exact path="/auth/login" element={<LoginPage />} />
            <Route exact path="/logout" element={<LogoutPage />} />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </BrowserRouter>
      </div>
  );
}

export default App;
