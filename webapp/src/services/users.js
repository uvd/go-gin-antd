import request from '../utils/request';

export async function query() {
  return request('/api/users');
}


//user 相关的

export async function list() {
  return request('/api/user/list');
}
export async function login(params) {
  return request('/api/user/login',{
      method :"post",
      data:params,
  });
}

export async function regist(params) {
  return request('/api/user/regist',{
      method :"post",
      data:params,
  });
}

export async function update(params) {
  return request('/api/user/update',{
      method :"post",
      data:params,
  });
}


