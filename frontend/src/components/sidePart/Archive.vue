<template>
  <div class="widget widget-archives">
    <h3 class="widget-title">归档</h3>
    <ul>
      <li v-for="(archive,index) in archiveList" :key="index">
        <a :href="'/archive/'+archive.Year+'/'+archive.Month"  >{{ archive.Pub_time }}（{{ archive.Cnt}}）</a>
      </li>
    </ul>
  </div>
</template>

<script>
import service from "@/utils/request"
export default {
    name:"Archive",
    data(){
      return {
        archiveList:[],
      }
    },
    methods:{
      showArchives:function(){
        service.get("/archive")
        .then(res=>{
          // console.log(res)
          var archiveList = res.data.data.archiveList
          var i 
          for(i=0;i<archiveList.length;i++){
            var object = archiveList[i]
            object.Year = object.Pub_time.slice(0,4)
            object.Month = object.Pub_time.slice(5,-1)
            this.archiveList.push(object)
          }
          // console.log(this.archiveList)
        })
      }
    },
    created(){
      this.showArchives()
    },
    
}
</script>