import { Module, MutationTree, GetterTree, ActionTree } from 'vuex'
import { RootState } from '~/store/index'

export interface BoardState {
  boardData: Array<any>
}
export const state: () => BoardState = (): BoardState => ({
  boardData: [
    {
      id: 1,
      title: '今週やること',
      position: 1,
      card: [
        {
          id: 1,
          title: '牛乳を買いに行くんじゃあああああああああああああああああああああああああああああああああああああ',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 1
        },
        {
          id: 2,
          title: '牛を狩に行く',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 2
        }
      ]
    },
    {
      id: 2,
      title: 'TODO',
      position: 2,
      card: [
        {
          id: 5,
          title: 'ミルクゲット',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 1
        },
        {
          id: 3,
          title: 'ミートゲット',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 2
        }
      ]
    }
  ]
})
const getters: GetterTree<BoardState, RootState> = {
  boardData(state: BoardState) {
    return state.boardData
  },
  kanbanData: (state: BoardState) => (index: number) => {
    return state.boardData[index].card
  }
}
const mutations: MutationTree<BoardState> = {
  updateBoardData(state: BoardState, newBoardData: any): void {
    state.boardData = newBoardData
  },
  updateKanbanData(state: BoardState, payload: any) {
    state.boardData[payload.index].card = payload.value
  }
}
export const Board: Module<BoardState, RootState> = {
  namespaced: true,
  state,
  getters,
  // actions,
  mutations
}
export default Board
