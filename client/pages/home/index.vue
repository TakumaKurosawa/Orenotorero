<template>
  <div>
    <h1>ユーザーTOPページ</h1>
    <MenuList></MenuList>
    <PersonalBoardList :card-titles="cardTitles"></PersonalBoardList>
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
  cardTitles = [
    { title: 'test1' },
    { title: 'test2' },
    { title: 'test3' },
    { title: 'test4' },
    { title: 'test5' },
    { title: 'test6' },
    { title: 'test7' },
    { title: 'test8' }
  ]
}
</script>
