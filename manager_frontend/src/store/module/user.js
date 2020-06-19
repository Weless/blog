import storageService from '@/service/storageService'
import userService from '@/service/userService'

const userModule = {
    namespaced: true,
    state: {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: JSON.parse(storageService.get(storageService.USER_INFO)), // 需要将字符串反序列化
        // userInfo:storageService.get(storageService.USER_INFO)
    },

    mutations: { // 状态只能在mutations中进行
        SET_TOKEN(state, token) {
            // 更新本地缓存
            storageService.set(storageService.USER_TOKEN, token)
            // 更新state
            state.token = token;
        },
        SET_USERINFO(state, userInfo) {
            storageService.set(storageService.USER_INFO, JSON.stringify(userInfo))
            state.userInfo = userInfo
        }
    },
    actions: {
        register(context, {
            name,
            password
        }) {
            return new Promise((resolve, reject) => {
                userService.register({name,password})
                    .then(res => {
                        // 保存token
                        context.commit('SET_TOKEN',res.data.data.token);
                        // storageService.set(storageService.USER_TOKEN, res.data.data.token);
                        return userService.info();
                    })
                    .then(res => {
                        // 保存用户信息
                        console.log(res.data.data.user)
                        context.commit('SET_USERINFO',res.data.data.user)
                        resolve(res);
                    }).catch((err) => {
                        reject(err);
                    })
            })
        },
        login(context, {
            name,
            password
        }) {
            return new Promise((resolve, reject) => {
                userService.login({name,password})
                    .then(res => {
                        // 保存token
                        context.commit('SET_TOKEN',res.data.data.token);
                        // storageService.set(storageService.USER_TOKEN, res.data.data.token);
                        return userService.info();
                    })
                    .then(res => {
                        // 保存用户信息
                        console.log(res.data.data.user)
                        context.commit('SET_USERINFO',res.data.data.user)
                        resolve(res);
                    }).catch((err) => {
                        reject(err);
                    })
            })
        },
        logout({commit}){
            // 清除token
            commit('SET_TOKEN','')
            storageService.set(storageService.USER_TOKEN,'')
            // 清除用户信息
            commit('SET_USERINFO','')
            storageService.set(storageService.USER_INFO,'')
        }
    }
}

export default userModule;
