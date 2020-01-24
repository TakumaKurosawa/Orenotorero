import { Module, MutationTree, GetterTree } from 'vuex'
import boardData from '../assets/mock/board.json'
import { RootState } from '~/store/index'

export interface BoardState {
  boardData: object
}
export const state: () => BoardState = (): BoardState => ({
  boardData
})
const getters: GetterTree<BoardState, RootState> = {
  getBoardData(state: BoardState) {
    return state.boardData
  }
}
const mutations: MutationTree<BoardState> = {
  setBoardData(state: BoardState, boardData: any): void {
    state.boardData = boardData
  }
}
// const actions: ActionTree<BoardState, RootState> = {
// //   nuxtServerInit ({ commit }, { req }) {
// //   if (req.headers.cookie) {
// //     const cookieData = cookieparser.parse(req.headers.cookie);
// //     const token = cookieData.token;
// //     const tokenAdmin = cookieData.tokenAdmin;
// //     if (tokenAdmin) {
// //       commit('admin/auth/loginAdmin', tokenAdmin);
// //     }
// //     if (token) {
// //       commit('login', token);
// //     }
// //   }
// // },
// // async setBoardData ({ commit }): Promise<any> {
// //   const boardData: any = this.getters.getboardData;
// // if (Object.keys(boardData).length === 0) {
// //   await this.$axios.get('/profile/')
// //     .then((res) => {
// //       console.log(res.data);
// //       commit('setUserData', res.data);
// //     }).catch((err) => {
// //       console.log(err);
// //     });
// // }
// }
// };
export const Board: Module<BoardState, RootState> = {
  namespaced: true,
  state,
  getters,
  // actions,
  mutations
}
export default Board
