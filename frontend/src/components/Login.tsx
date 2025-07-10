import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom'; 
import API from '../../utils/api';

const Login: React.FC = () => {
  const [form, setForm] = useState({ email: '', password: '' });
  const [message, setMessage] = useState('');
  const [error, setError] = useState(false);
  const navigate = useNavigate(); // 👈 Import the hook

const handleSubmit = async (e: React.FormEvent) => {
  e.preventDefault();
  setError(false);
  setMessage('');
  try {
    const response = await API.post('/login', form);

    // ✅ Save token to localStorage
    const token = response.data.token;
    localStorage.setItem('token', token);

    setMessage('Login successful!');

    // ✅ Redirect after a short delay
    setTimeout(() => {
      navigate('/list');
    }, 500);
  } catch (err: any) {
    setError(true);
    setMessage(`Login failed: ${err.response?.data?.error || 'Unknown error'}`);
  }
};


  return (
    <div className="max-w-md mx-auto mt-16 p-8 bg-white rounded-lg shadow-lg">
      <h2 className="text-3xl font-semibold text-center mb-6">Login</h2>
      
      <form onSubmit={handleSubmit} className="space-y-5">
        <input
          type="email"
          placeholder="Email"
          className="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          value={form.email}
          onChange={(e) => setForm({ ...form, email: e.target.value })}
          required
        />
        <input
          type="password"
          placeholder="Password"
          className="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          value={form.password}
          onChange={(e) => setForm({ ...form, password: e.target.value })}
          required
        />
        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 transition-colors"
        >
          Login
        </button>
      </form>

      {message && (
        <p
          className={`mt-4 text-center text-sm ${
            error ? 'text-red-600' : 'text-green-600'
          }`}
        >
          {message}
        </p>
      )}

      <div className="mt-6 text-center text-sm text-gray-600">
        Don’t have an account?{' '}
        <Link to="/register" className="text-blue-600 hover:underline">
          Register here
        </Link>
      </div>
    </div>
  );
};

export default Login;
