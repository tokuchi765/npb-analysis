import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

const rest = (() => {
  const client = axios.create({
    baseURL: 'http://localhost:8081',
    timeout: 15000,
  });
  return {
    client,
    get: <T = any, R = AxiosResponse<T>>(url: string): Promise<R> => {
      return client.get(url);
    },
    getParams: <T = any, R = AxiosResponse<T>>(
      url: string,
      params: AxiosRequestConfig
    ): Promise<R> => {
      return client.get(url, params);
    },
  };
})();

export { rest };
