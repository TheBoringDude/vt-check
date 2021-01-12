<template>
  <!-- full screen results dialog -->
  <v-dialog
    v-model="show_results"
    fullscreen
    hide-overlay
    transition="dialog-bottom-transition"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn color="secondary" @click="$emit('reset-new')">
        Check Another File
      </v-btn>
      <v-btn color="primary" dark v-bind="attrs" v-on="on">
        Show Results
      </v-btn>
    </template>
    <v-card>
      <v-toolbar dark color="primary">
        <v-btn icon dark @click="show_results = hideResults">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>Results: {{ upload_filename }}</v-toolbar-title>
        <!-- <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn dark text @click="show_results = false">Save</v-btn>
        </v-toolbar-items> -->
      </v-toolbar>
      <v-list three-line subheader>
        <v-subheader>General</v-subheader>
        <v-container>
          <v-row dense>
            <v-col
              cols="4"
              v-for="result in data.data.attributes.results"
              v-bind:key="result.engine_name"
            >
              <v-card>
                <v-list-item three-line>
                  <v-list-item-content>
                    <div class="overline mb-2 text-truncate">
                      {{ result.engine_name }}
                    </div>
                    <v-list-item-title class="headline mb-1">
                      {{ result.category }}
                    </v-list-item-title>
                    <v-list-item-subtitle class="result-subtitle">
                      {{ result.result }}
                    </v-list-item-subtitle>
                  </v-list-item-content>

                  <v-list-item-avatar
                    tile
                    size="30"
                    :color="handleResultColor(result.category)"
                  >
                    <!-- undetected -->
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 20 20"
                      fill="#ffffff"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                        clip-rule="evenodd"
                      />
                    </svg>
                    <!-- unsupported-type || failure -->
                    <!-- <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 20 20"
                      fill="#ffffff"
                      v-else-if="
                        result.category === 'type-unsupported' ||
                        result.category === 'failure'
                      "
                    >
                      <path
                        fill-rule="evenodd"
                        d="M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z"
                        clip-rule="evenodd"
                      />
                    </svg> -->
                    <!-- something else, error maybe -->
                    <!-- <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 20 20"
                      fill="#ffffff"
                      v-else
                    >
                      <path
                        fill-rule="evenodd"
                        d="M10 1.944A11.954 11.954 0 012.166 5C2.056 5.649 2 6.319 2 7c0 5.225 3.34 9.67 8 11.317C14.66 16.67 18 12.225 18 7c0-.682-.057-1.35-.166-2.001A11.954 11.954 0 0110 1.944zM11 14a1 1 0 11-2 0 1 1 0 012 0zm0-7a1 1 0 10-2 0v3a1 1 0 102 0V7z"
                        clip-rule="evenodd"
                      />
                    </svg> -->
                  </v-list-item-avatar>
                </v-list-item>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-list>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'Result',
  computed: {
    hideResults() {
      return false
    },
  },
  data() {
    return {
      show_results: false,
    }
  },
  methods: {
    handleResultColor(resCategory) {
      if (resCategory == 'undetected') {
        return 'green accent-3'
      }

      if (
        resCategory == 'type-unsupported' ||
        resCategory == 'failure' ||
        resCategory == 'timeout'
      ) {
        return 'grey lighten-1'
      }

      return 'red accent-2'
    },
  },
  props: {
    upload_filename: String,
    data: Object,
  },
}
</script>

<style scoped>
.result-subtitle {
  font-size: 1.2em;
}
</style>
