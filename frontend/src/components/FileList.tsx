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

    API
      .get("/getFiles", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => setFiles(res.data))
      .catch((err) =>
        setError(` Failed to fetch files: ${err.response?.data?.error || "Unknown error"}`)
      );
  }, []);

  return (
    <div className="max-w-4xl mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl font-bold mb-4 text-center">📁 Your Files</h2>

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
              <td className="p-2 border">{new Date(file.uploadedAt).toLocaleString()}</td>
              <td className="p-2 border">
                <a
                  href={`http://localhost:5000/${file.filePath}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-blue-600 hover:underline"
                >
                  Download
                </a>
              </td>
              <td className="p-2 border">
                {file.isShared ? (
                  <a
                    href={`http://localhost:5000/shared/${file.shareToken}`}
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
