import { Module, MutationTree, GetterTree, ActionTree } from 'vuex'
import { RootState } from '~/store/index'

export interface MyBoardsState {
  myBoardsData: Array<any>
}

export const state: () => MyBoardsState = (): MyBoardsState => ({
  myBoardsData: []
})

const getters: GetterTree<MyBoardsState, RootState> = {
  myBoardsData(state: MyBoardsState) {
    return state.myBoardsData
  }
}
const mutations: MutationTree<MyBoardsState> = {
  updateMyBoardsData(state: MyBoardsState, newMyBoardsData: any): void {
    state.myBoardsData = newMyBoardsData
  }
}
const actions: ActionTree<MyBoardsState, RootState> = {
  async fetchMyBoardsData({ rootGetters, commit }: any): Promise<any> {
    await this.$axios
      .$get('/board', {
        headers: {
          Authorization: 'Bearer ' + rootGetters['auth/getAuthToken']
        },
        data: {}
      })
      .then((res: any) => {
        console.log(res.data)
        commit('updateMyBoardsData', res.data)
      })
      .catch((err: any) => {
        console.log(err)
      })
  },
  async createBoard(
    { rootGetters, dispatch }: any,
    payload: any
  ): Promise<any> {
    await this.$axios
      .$post(
        '/board',
        { title: payload.title, img: payload.img },
        {
          headers: {
            Authorization: 'Bearer ' + rootGetters['auth/getAuthToken']
          }
        }
      )
      .then((res: any) => {
        console.log(res.data)
        dispatch('fetchMyBoardsData')
      })
      .catch((err: any) => {
        console.log(err)
      })
  }
}
export const MyBoards: Module<MyBoardsState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
export default MyBoards
