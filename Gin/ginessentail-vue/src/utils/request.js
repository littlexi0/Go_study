import axios from 'axios';
import storageService from '@/service/storageService';

const server = axios.create({
    // baseURL: process.env.VUE_APP_BASE_API,
    baseURL: 'http://localhost:8081/api/',
    timeout: 1000*5,
    // headers: { Authorization: "Bearer "+storageService.get(storageService.USER_TOKEN)},
}); 

server.interceptors.request.use(function (config){
    Object.assign(config.headers, { Authorization: "Bearer "+storageService.get(storageService.USER_TOKEN)});
    return config;
},(error) =>{
    return Promise.reject(error);
});


export default server;