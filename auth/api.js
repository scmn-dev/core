import axios from "axios";

export default axios.create({
  baseURL: "$PASSWORD_MANAGER_API_WEBSITE_URL",
  responseType: "json",
});
