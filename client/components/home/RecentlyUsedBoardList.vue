<template>
  <div>
    <LabelIcon :text-label="'最近の表示'" :icon-name="'mdi-clock'"></LabelIcon>
    <v-row>
      <Card
        v-for="item in cardCount"
        :key="item.id"
        :card-title="item.name"
        :card-height="100"
        :card-width="180"
        :card-img="item.img"
      ></Card>
    </v-row>
  </div>
</template>
<script lang="ts">
import { Vue, Prop, Component } from 'nuxt-property-decorator'
import LabelIcon from '@/components/atom/LabelIcon.vue'
import Card from '@/components/atom/Card.vue'

@Component({
  components: {
    LabelIcon,
    Card
  }
})
export default class RecentlyUsedBoardList extends Vue {
  get cardCount() {
    return this.boardData.slice(0, 4)
  }

  @Prop({ type: Array, required: true })
  boardData!: Array<string>
}
</script>
