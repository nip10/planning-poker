import Axios, { AxiosRequestConfig } from "axios";
import Router from "next/router";

export const urls = {
  test: `http://localhost:3334`,
  development: "http://localhost:3001",
  production: "https://api.idontknowyet.fooooo",
};

const api = Axios.create({
  baseURL: urls[process.env.NODE_ENV],
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

export default api;
