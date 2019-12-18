<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <PersonalBoardList :board-data="boardData"></PersonalBoardList>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import MenuList from '../../components/home/MenuList.vue'
import PersonalBoardList from '../../components/home/PersonalBoardList.vue'

@Component({
  components: {
    MenuList,
    PersonalBoardList
  }
})
export default class Home extends Vue {
  boardData = [
    {
      title: 'test1',
      img: 'https://pro.foto.ne.jp/free/img/images_thumb/yam0065-049.jpg'
    },
    { title: 'test2', img: 'https://cdn.vuetifyjs.com/images/cards/docks.jpg' },
    {
      title: 'test3',
      img: 'https://pro.foto.ne.jp/share/v5/linkthumb_stock/img/36/49002993.jpg'
    },
    {
      title: 'test4',
      img: 'https://pro.foto.ne.jp/share/v5/linkthumb_stock/img/36/6918622.jpg'
    },
    {
      title: 'test5',
      img: 'https://pro.foto.ne.jp/share/v5/linkthumb_stock/img/36/35242924.jpg'
    },
    {
      title: 'test6',
      img: 'https://pro.foto.ne.jp/share/v5/linkthumb_stock/img/36/36758373.jpg'
    },
    {
      title: 'test7',
      img: 'https://pro.foto.ne.jp/share/v5/linkthumb_stock/img/36/639385.jpg'
    },
    {
      title: 'test8',
      img: 'https://pro.foto.ne.jp/share/v5/linkthumb_stock/img/37/27330761.jpg'
    }
  ]
}
</script>
