<template>
  <svg
    id="svg8"
    xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:cc="http://creativecommons.org/ns#"
    xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
    xmlns:svg="http://www.w3.org/2000/svg"
    xmlns="http://www.w3.org/2000/svg"
    xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
    xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
    width="48.30785mm"
    height="48.107243mm"
    viewBox="0 0 48.30785 48.107242"
    version="1.1"
    sodipodi:docname="logo.svg"
    inkscape:version="0.92.4 5da689c313, 2019-01-14"
  >
    <defs id="defs2" />
    <g
      id="layer1"
      inkscape:label="Layer 1"
      inkscape:groupmode="layer"
      style="display:inline"
      transform="translate(-61.874306,-105.71226)"
    >
      <path
        id="rect815"
        sodipodi:type="inkscape:offset"
        inkscape:radius="1.4652326"
        inkscape:original="M 90.521484 107.59961 L 63.761719 133.72266 L 81.017578 134.6582 L 81.535156 151.93164 L 108.29688 125.80859 L 90.521484 107.59961 z "
        style="fill:#59344f;fill-opacity:1;stroke:#59344f;stroke-width:0.84500003;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
        :style="`fill:${fillColor};stroke:${fillColor}`"
        d="m 90.517578,106.13477 a 1.4653791,1.4653791 0 0 0 -1.019531,0.41601 l -26.759766,26.12305 a 1.4653791,1.4653791 0 0 0 0.94336,2.51172 l 15.912109,0.86328 0.476562,15.92578 a 1.4653791,1.4653791 0 0 0 2.488282,1.00586 l 26.761716,-26.12305 a 1.4653791,1.4653791 0 0 0 0.0254,-2.07226 L 91.570312,106.57617 a 1.4653791,1.4653791 0 0 0 -1.052734,-0.4414 z"
      />
      <path
        id="path856"
        style="fill-opacity:1;stroke-width:0.04370949;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
        d="m 77.184667,121.84485 16.21664,0.90728 0.51546,16.23382 -3.6393,3.47272 0.002,-16.73938 -16.73456,-0.40172 z"
        :style="`fill:${innerColor};stroke:${fillColor}`"
        inkscape:connector-curvature="0"
        sodipodi:nodetypes="ccccccc"
      />
      <path
        id="path839"
        sodipodi:nodetypes="ccccccc"
        inkscape:connector-curvature="0"
        d="m 72.632527,126.11897 16.21664,0.90728 0.51545,16.23382 -3.63929,3.47272 0.002,-16.73938 -16.73456,-0.40173 z"
        style="fill-opacity:1;stroke-width:0.04370949;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
        :style="`fill:${innerColor};stroke:${fillColor}`"
      />
      <rect
        id="rect866"
        style="fill-opacity:1;fill-rule:nonzero;stroke-width:0.04504975;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
        :style="`fill:${innerColor};stroke:${fillColor}`"
        width="1.8183577"
        height="14.304414"
        x="162.43921"
        y="11.314446"
        transform="rotate(45.691794)"
      />
    </g>
  </svg>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BrandIcon",
  props: {
    fillColor: {
      type: String,
      default: "#59344f",
    },
    innerColor: {
      type: String,
      default: "#ffffff",
    },
  },
});
</script>
