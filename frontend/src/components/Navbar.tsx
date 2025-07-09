import { Link, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

const Navbar = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    setLoggedIn(!!token);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  if (!loggedIn) return null;

  return (
    <nav className="bg-gray-800 text-white px-6 py-3 flex justify-between items-center">
      <div className="text-lg font-bold">File Manager</div>
      <div className="space-x-4">
        <Link to="/upload" className="hover:underline">
          Upload
        </Link>
        <Link to="/list" className="hover:underline">
          File List
        </Link>
        <button onClick={handleLogout} className="hover:underline">
          Logout
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
