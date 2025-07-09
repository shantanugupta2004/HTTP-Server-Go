import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import API from "../../utils/api";

const FileUpload = () => {
  const [file, setFile] = useState<File | null>(null);
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/login");
    }
  }, [navigate]);

  const handleUpload = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!file) {
      setMessage("⚠️ Please select a file to upload.");
      return;
    }

    const token = localStorage.getItem("token");
    if (!token) {
      setMessage("❌ You must be logged in to upload files.");
      navigate("/login");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      await API.post("/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
          "Authorization": `Bearer ${token}`,
        },
      });

      setMessage("✅ File uploaded successfully!");
      setFile(null);
    } catch (err: any) {
      setMessage(`❌ Upload failed: ${err.response?.data?.error || "Unknown error"}`);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-xl font-bold mb-4 text-center">📤 Upload File</h2>
      <form onSubmit={handleUpload} className="space-y-4">
        <input
          type="file"
          onChange={(e) => setFile(e.target.files?.[0] || null)}
          className="w-full border px-4 py-2 rounded"
        />
        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
        >
          Upload
        </button>
      </form>
      {message && <p className="mt-4 text-sm text-center">{message}</p>}
    </div>
  );
};

export default FileUpload;
