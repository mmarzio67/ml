import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:5900",
});

//instance.defaults.headers.common["SOMETHING"] = "something";

export default instance;
