import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import FileUpload from "./components/FileUpload";
import FileList from "./components/FileList";
import AuthRoute from "./components/AuthRoute";
import Navbar from "./components/Navbar";

function App() {
  const token = localStorage.getItem("token");

  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Navigate to={token ? "/list" : "/login"} replace />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/upload" element={<AuthRoute><FileUpload /></AuthRoute>} />
        <Route path="/list" element={<AuthRoute><FileList /></AuthRoute>} />
      </Routes>
    </Router>
  );
}

export default App;
