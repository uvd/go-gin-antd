import {list,regist,login,update} from "../services/users"


export default{
    
    //命名空间
    namespace:"user",
    //状态
    state:{
        user:{}

    },
   //这个函数干嘛的  还没有理解哦
    subscriptions:{
       //大概就是进来的时候默认触发的
        setup({}){

        }
    },

   //触发的地方
    effects:{
        *login({payload},{call,put}){
            const data = yield call(login,payload)

            if(data){
                yield put({
                    type:'loginSuccess',
                    payload:{
                        user:data.data
                    }
                })
            }
        },

    },
   
   //定义触发的类型
    reducers:{
        loginSuccess(state){
            return {...state}
        }
    }
} 