import { useEffect, useState } from "react";
import axios from "axios";
import "./App.css";

type Response = {
  message: string;
};

function App() {
  const [message, setMessage] = useState<string>("");

  useEffect(() => {
    axios
      .get<Response>("http://localhost:8000/api/hello")
      .then((response) => {
        console.log("Response data:", response.data);
        setMessage(response.data.message);
      })
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <div>
      <h1> hi </h1>
      <h1> {message} </h1>
    </div>
  );
}

export default App;
