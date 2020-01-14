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
    id: 0,
    img: '',
    name: ''
  },
  authToken: ''
})
const getters: GetterTree<AuthState, RootState> = {
  getAuthedUser(state: AuthState) {
    return state.authedUser
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
  login({ commit }, email: string, pass: string): void {
    const res = this.$axios.$post('/login', {
      email,
      password: pass
    })
    console.log(res.data)
    commit('updateIsAuth', true)
  },
  test(): void {
    const res = this.$axios.$get('/ping')
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
