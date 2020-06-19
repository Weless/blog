import request from "@/utils/request";

const register = ({ name, password }) => {
  return request.post("auth/register", { name, password });
};

const info = () => {
  return request.get("auth/info");
};

const login = ({ name, password }) => {
  return request.post("auth/login", { name, password });
};

export default {
  register,
  info,
  login
};
