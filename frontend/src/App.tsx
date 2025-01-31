import { useState } from 'react';
import './App.css';

const baseURL = import.meta.env.VITE_BACKEND_API;

function App() {
  const [response, setResponse] = useState<Ping | null>(null);

  const ping = async () => {
    const res = await fetch(baseURL + '/api/ping');
    if (res.ok) {
      const json = await res.json();
      const time = new Date(json.time).toUTCString();
      setResponse({ message: json.message, time });
    }
  };

  return (
    <>
      <h1>Bank Analytics</h1>
      <button onClick={ping}>Ping</button>
      {response && (
        <div>
          {response.message} - {response.time}
        </div>
      )}
    </>
  );
}

export default App;

type Ping = {
  message: string;
  time: string;
};
