import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import FileUpload from "./components/FileUpload";
import FileList from "./components/FileList";

function App(){
  return(
    <Router>
      <Routes>
        <Route path="/login" element={<Login/>}/>
        <Route path="/register" element={<Register/>}/>
        <Route path="/upload" element={<FileUpload/>}/>
        <Route path="/list" element={<FileList/>}/>
      </Routes>
    </Router>
  )
}

export default App;