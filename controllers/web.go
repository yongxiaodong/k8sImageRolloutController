package controllers

import "github.com/gin-gonic/gin"

func WebHtml(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, `<!DOCTYPE html>
<html>
  <head>
    <title>Ó¦ÓÃ¿ìËÙ»Ø¹ö</title>
    <link rel="stylesheet" href="https://unpkg.com/element-ui@2.13.2/lib/theme-chalk/index.css">
    <script src="https://cdn.jsdelivr.net/npm/vue@2.6.11/dist/vue.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
    <script src="https://unpkg.com/element-ui@2.13.2/lib/index.js"></script>
    <style>
      .center
      {
        margin: auto;
        width: 100%;
        padding: 70px 0;
        /* border: 3px solid green;
        padding: 10px */
      } 
      .center2
      {
        margin: left;
        width: 100%;
        padding: 70px 0;
        /* border: 3px solid green;
        padding: 10px */
      }
      .center3 {

      }
      .el-row {
    margin-bottom: 20px;

  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
      </style>
  </head>
  <body>

    <div >
      <div id="app">
        <el-row>
          <el-col :span="24"><div class="grid-content bg-purple-dark"></div></el-col>
        </el-row>

        <el-container style="height: 500px; border: 1px solid #eee">
          
          <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
            <el-menu :default-openeds="['1', '3']">
              <el-submenu index="1">
                <template slot="title"><i class="el-icon-message"></i>µ¼º½</template>

                <el-submenu index="1-1">
                  <template slot="title">Ó¦ÓÃ²Ù×÷</template>
                  <el-menu-item index="1-4-1">»Ø¹ö</el-menu-item>
                </el-submenu>
              </el-submenu>
              
  
            </el-menu>
          </el-aside>
          <!-- ÓÒ²à¹¦ÄÜ -->
          <div >
          
          <el-container>
            
  
            <el-main>

              <el-select v-model="selectedOption1" filterable placeholder="ÇëÑ¡ÔñÏîÄ¿(¿ÉËÑË÷)">
                <el-option v-for="option in options1" :key="option" :label="option" :value="option"></el-option>
              </el-select>
      
              <el-select v-model="selectedOption2" v-if="options2.length" filterable placeholder="ÇëÑ¡Ôñ»Ø¹ö°æ±¾">
                <el-option v-for="option in options2" :key="option.Image" :label="'¾µÏñ:' + option.Image + '  ·¢²¼Ê±¼ä:' + option.CreatedAt" :value="option"></el-option>
              </el-select>
      
            
                <el-button  @click="dialogVisible = true" type="danger" :disabled="selectedOption1 === '' || selectedOption2 === ''">Ìá½»»Ø¹ö</el-button>
                <el-dialog
                  title="¾¯¸æ"
                  :visible.sync="dialogVisible"
                  width="30%"
                  :before-close="handleClose">
                  <span>Ö±½ÓÌæ»»ÏßÉÏÓ¦ÓÃ°æ±¾£¬Çë½÷É÷Ìá½»!!!</span>
                  <span slot="footer" class="dialog-footer">
                    <el-button @click="dialogVisible = false">È¡ Ïû</el-button>
                    <el-button :plain="true" type="primary" @click="submit">È·ÈÏ»Ø¹ö</el-button>
                  </span>
                </el-dialog>
                <el-button @click="ref" type="primary">Ë¢ÐÂ°æ±¾</el-button>
      
                <div style="height: 60px;"></div>
      
                <!-- black box to display selectedOption1 value -->
                <div class="center2">
                  <el-card class="box-card">
                    <div slot="header" class="clearfix">
                      <span>µ±Ç°Ñ¡ÔñÄÚÈÝÐÅÏ¢</span>
                    </div>
                    <pre v-if="selectedOption2">
                      ÏîÄ¿£º{{ selectedOption1 }}
                      ·¢²¼Ê±¼ä£º{{ selectedOption2.CreatedAt }}
                      Ä¿±êCommit ID£º{{ selectedOption2.Image.split(':')[selectedOption2.Image.split(':').length - 1] }}
                      ¾µÏñÃû³Æ: {{selectedOption2.Image}}
                    </pre>
                  </el-card>
                </div>

            </el-main>
          </el-container>
        </div>
        <!-- ÓÒ²à¹¦ÄÜ½áÊø -->
        </el-container>



        
      </div>
  </div>
    <script>
      new Vue({
        el: '#app',
        data() {
          return {
            options1: [],
            options2: [],
            selectedOption1: '',
            selectedOption2: '',
            dialogVisible: false
          };
        },
        created() {
          axios.get('http://192.168.104.120:16100/api/v1/getAllDp')
            .then(response => {
              this.options1 = response.data.map(item => item.DpName);
            })
            .catch(error => {
              console.error(error);
            });
        },
        watch: {
          selectedOption1(newVal) {
            this.newVal = newVal
            axios.get("http://192.168.104.120:16100/api/v1/getDpImage?name=${newVal}")
              .then(response => {
                this.options2 = response.data;
              })
              .catch(error => {
                console.error(error);
              });
          }
        },
        methods: {
          submit() {
            this.dialogVisible = false;
            const dpName = this.selectedOption1;
            const selectedImage = this.selectedOption2;
    
            if (!dpName || !selectedImage) {
              return;
            }
    
            axios.post('http://192.168.104.120:16100/api/v1/commit', {
              dpName,
              selectedImage
            })
              .then(response => {


                if (response.data.Status == 0) {
                  this.$message({
                  message: "³É¹¦" + response.data.Message,
                  type: 'success'
                });
                } else {
                  this.$message({
                  message: "´íÎó" + response.data.Message,
                  type: 'warning'
                });
              }})
              .catch(error => {
                console.error(error);
                this.$message.error('·¢ÉúÒì³£');
              });
          },
          ref() {
            axios.get("http://192.168.104.120:16100/api/v1/getDpImage?name=${this.newVal}")
              .then(response => {
                this.options2 = response.data;
              })
              .catch(error => {
                console.error(error);
              });
          },
          handleClose(done) {
            this.$confirm('È·ÈÏ¹Ø±Õ£¿')
              .then(_ => {
                done();
              })
              .catch(_ => {});
          }
    

        }
      });
    </script>
    </body>
    </html>`)

}
