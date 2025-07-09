import { useEffect, useState } from "react";
import API from "../../utils/api";

type FileEntry = {
  id: number;
  fileName: string;
  filePath: string;
  fileSize: number;
  uploadedAt: string;
  isShared: boolean;
  shareToken: string;
  userID: number;
};

const formatSize = (size?: number) => {
  if (typeof size !== "number" || isNaN(size)) return "Unknown";
  const units = ["B", "KB", "MB", "GB"];
  let i = 0;
  while (size >= 1024 && i < units.length - 1) {
    size /= 1024;
    i++;
  }
  return `${size.toFixed(2)} ${units[i]}`;
};

const FileList = () => {
  const [files, setFiles] = useState<FileEntry[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("You must be logged in to view files.");
      return;
    }

    API.get("/getFiles", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => setFiles(res.data))
      .catch((err) =>
        setError(
          `Failed to fetch files: ${
            err.response?.data?.error || "Unknown error"
          }`
        )
      );
  }, []);

  const handleDownload = async (fileName: string) => {
    const token = localStorage.getItem("token");
    if (!token) {
      alert("You must be logged in to download files.");
      return;
    }

    try {
      const response = await fetch(
        `http://localhost:5000/downloadFile?name=${encodeURIComponent(fileName)}`,
        {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      if (!response.ok) throw new Error("Download failed");

      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = fileName;
      document.body.appendChild(a);
      a.click();
      a.remove();
      window.URL.revokeObjectURL(url);
    } catch (err) {
      alert("Error downloading file");
      console.error(err);
    }
  };

  return (
    <div className="max-w-4xl mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl font-bold mb-4 text-center">Your Files</h2>

      {error && <p className="text-red-600 text-center">{error}</p>}

      <table className="w-full table-auto mt-4 border">
        <thead>
          <tr className="bg-gray-100">
            <th className="p-2 border">File Name</th>
            <th className="p-2 border">Size</th>
            <th className="p-2 border">Uploaded</th>
            <th className="p-2 border">Download</th>
            <th className="p-2 border">Share Link</th>
          </tr>
        </thead>
        <tbody>
          {files.map((file) => (
            <tr key={file.id} className="text-center">
              <td className="p-2 border">{file.fileName}</td>
              <td className="p-2 border">{formatSize(file.fileSize)}</td>
              <td className="p-2 border">
                {new Date(file.uploadedAt).toLocaleString()}
              </td>
              <td className="p-2 border">
                <button
                  onClick={() => handleDownload(file.fileName)}
                  className="text-blue-600 hover:underline"
                >
                  Download
                </button>
              </td>
              <td className="p-2 border">
                {file.isShared ? (
                  <a
                    href={`http://localhost:5000/share/${file.shareToken}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-green-600 hover:underline"
                  >
                    View Link
                  </a>
                ) : (
                  <span className="text-gray-500">Private</span>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default FileList;
