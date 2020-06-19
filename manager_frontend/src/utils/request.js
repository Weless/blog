import axios from "axios";
// import { Message } from "element-ui";
// import { getToken, getUsername } from "@/utils/app";
import storageService from "../service/storageService";
// import Router from "@/router";

// const BASEURL = process.env.NODE_ENV === "production" ? "" : "/devApi";
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 5000
});

service.interceptors.request.use(
  function(config) {
    Object.assign(config.headers, {
      Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}`
    });
    return config;
  },
  function(error) {
    return Promise.reject(error);
  }
);

// service.interceptors.response.use(
//   function(response) {
//     let data = response.data;
//     if (data.resCode !== 0) {
//       Message.error(data.resMsg);
//       if (data.resCode === 100) {
//         Router.push({ path: "/login" });
//       }
//       return Promise.reject(response);
//     } else {
//       return response;
//     }
//   },
//   function(error) {
//     return Promise.reject(error);
//   }
// );

export default service;
