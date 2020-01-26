import { Module, MutationTree, GetterTree, ActionTree } from 'vuex'
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
  // cardData: (state: BoardState) => (index: number) => {
  //   return state.boardData[index].card
  // }
}
const mutations: MutationTree<BoardState> = {
  updateBoardData(state: BoardState, newBoardData: Array<object>): void {
    state.boardData = newBoardData
  }
}
const actions: ActionTree<BoardState, RootState> = {
  async fetchBoardData({ rootState, commit }, boardId): Promise<any> {
    // axios -> データを取ってくる
    const boardData: Array<object> = await this.$axios
      .$get('/kanban', {
        headers: {
          Authorization: 'Bearer ' + rootState.auth.authToken,
          BoardId: boardId
        },
        data: {}
      })
      .catch((err: any) => {
        console.log(err)
      })

    // mutationのupdateBoardDataを呼び出す
    commit('updateBoardData', boardData)
  }
}

export const Board: Module<BoardState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
export default Board
