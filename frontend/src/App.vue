<template>
  <v-app id="inspire">
    <v-main>
      <v-container fluid>
        <v-layout justify-center align-center class="main-layout">
          <main>
            <header>
              <h1>
                VirusTotal File Checker
              </h1>
            </header>

            <!-- display file and upload information -->
            <div class="file-info-upload" v-if="uploading">
              <p>
                Filename:
                <span class="filename">{{ upload_filename }}</span>
              </p>
              <hr />
              <div class="uploading">
                <p>{{ upload_status }}</p>
                <v-progress-linear
                  indeterminate
                  color="green"
                ></v-progress-linear>
              </div>
            </div>
            <!-- results button actons -->
            <div class="results-main" v-if="done_upload">
              <Result :results="result" :upload_filename="upload_filename" />
            </div>

            <!-- main form -->
            <div class="main-form" v-if="form_upload">
              <p>Please select a file to upload and check...</p>
              <v-btn x-large block primary @click="SelectFile">
                Choose a file...
              </v-btn>
            </div>

            <!-- footer -->
            <footer>
              <v-dialog v-model="dialog" persistent max-width="500">
                <template v-slot:activator="{ on, attrs }">
                  <button v-bind="attrs" v-on="on">
                    <v-icon x-large class="settings-btn">mdi-cogs</v-icon>
                  </button>
                </template>
                <v-card>
                  <v-card-title class="headline">
                    Your VirusTotal API Key
                  </v-card-title>
                  <v-container>
                    <v-col cols="12">
                      <v-text-field
                        label="Api Key*"
                        required
                        value="1"
                      ></v-text-field>
                    </v-col>
                  </v-container>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="green darken-1"
                      text
                      @click="dialog = false"
                      outlined
                    >
                      Save
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </footer>
          </main>
        </v-layout>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Result from './components/Results.vue'
export default {
  data: () => ({
    dialog: false,
    form_upload: false,
    uploading: false,
    upload_filename: '',
    upload_status: 'Uploading...',
    done_upload: true,
    result: {},
  }),
  watch: {
    upload_filename: {
      handler(upload_filename) {
        if (upload_filename === '') {
          return ''
        }
        this.uploading = true
        this.form_upload = false
      },
    },
  },
  methods: {
    SelectFile() {
      window.backend.File.SelectFileUpload().then((file) => {
        if (file !== '') {
          this.upload_filename = file
        }
      })
    },
  },
  components: {
    Result,
  },
}
</script>

<style>
.file-info-upload {
  padding: 2em;
  margin-top: 12px;
  border: 1px solid #e0e0e0;
}
.file-info-upload p {
  font-size: 1.2em;
}
.file-info-upload .filename {
  font-weight: bold;
  text-decoration: underline;
  color: #4caf50;
}
.file-info-upload .uploading {
  margin-top: 10px;
}
.file-info-upload .uploading p {
  font-size: 1em;
}
.file-info-upload hr {
  border: 1px solid #e0e0e0;
}
.results-main {
  text-align: center;
  margin-top: 14px;
}
.results-main button {
  margin: 0 1px 0 1px;
}
.main-layout {
  height: 550px;
  width: 900px;
  margin: 0px auto;
  position: relative;
}
header h1 {
  font-size: 4em;
  letter-spacing: 2px;
  text-align: center;
}
.main-form {
  margin-top: 10px;
}
footer button {
  position: absolute;
  bottom: 0px;
  right: 0px;
}
.settings-btn {
  height: 64px;
  background: #e0e0e0;
  width: 64px;
}
.settings-btn:hover {
  background: #bdbdbd;
}
</style>
