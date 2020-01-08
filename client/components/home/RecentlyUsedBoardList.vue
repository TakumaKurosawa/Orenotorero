<template>
  <div>
    <LabelIcon :text-label="'最近の表示'" :icon-name="'mdi-clock'"></LabelIcon>
    <v-row>
      <Card
        v-for="i in 4"
        :key="boardData[i - 1].id"
        :card-title="boardData[i - 1].name"
        :card-height="100"
        :card-width="180"
        :card-img="boardData[i - 1].img"
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
  @Prop({ type: Array, required: true })
  boardData!: Array<string>
}
</script>
