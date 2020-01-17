import { Module, MutationTree, GetterTree, ActionTree } from 'vuex'
export interface RootState {
  version: string
}
export interface AuthState {
  isAuth: boolean
  authedUser: object
  authToken: string
}
export const state: () => AuthState = (): AuthState => ({
  isAuth: false,
  authedUser: {
    img: '',
    name: ''
  },
  authToken: ''
})
const getters: GetterTree<AuthState, RootState> = {
  getAuthedUser(state: AuthState) {
    return state.authedUser
  },
  getIsAuth(state: AuthState) {
    return state.isAuth
  },
  getAuthToken(state: AuthState) {
    return state.authToken
  }
}
const mutations: MutationTree<AuthState> = {
  updateIsAuth(state: AuthState, isAuth: boolean): void {
    state.isAuth = isAuth
  },
  updateUserData(state: AuthState, user: object): void {
    state.authedUser = user
  },
  updateToken(state: AuthState, token: string): void {
    state.authToken = token
  }
}
const actions: ActionTree<AuthState, RootState> = {
  // login({ commit }, email: string, pass: string): void {
  //   const res = this.$axios.$post('/login', {
  //     email,
  //     password: pass
  //   })
  //   console.log(res.data)
  //   commit('updateIsAuth', true)
  // },
  async login({ commit }: any, payload: object): Promise<any> {
    await this.$axios
      .post('/login', {
        payload
      })
      .then((res: any) => {
        console.log(res.data)
        commit('updateIsAuth', true)
        commit('updateToken', res.data.token)
      })
      .catch((err: any) => {
        console.log(err)
      })
  },
  async signup({ commit }: any, payload: object): Promise<any> {
    await this.$axios
      .post('/user/create', payload)
      .then((res: any) => {
        console.log(res.data)
        commit('updateIsAuth', true)
        commit('updateToken', res.data.token)
      })
      .catch((err: any) => {
        console.log(err)
      })
  },
  async test(): Promise<any> {
    const res = await this.$axios.get('/ping')
    console.log(res.data)
  }
}
export const Auth: Module<AuthState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
export default Auth
