import { HMSVideoPluginType } from "@100mslive/hms-video";

export class GrayscalePlugin {
  getName() {
    return "grayscale-plugin";
  }

  isSupported() {
    return true;
  }

  async init() {}

  getPluginType() {
    return HMSVideoPluginType.TRANSFORM;
  }

  stop() {}

  /**
   * @param input {HTMLCanvasElement}
   * @param output {HTMLCanvasElement}
   * @param skipProcessing {boolean}
   */
  processVideoFrame(input, output, skipProcessing=true) {
    const width = input.width;
    const height = input.height;
    output.width = width;
    output.height = height;
    const inputCtx = input.getContext("2d");
    const outputCtx = output.getContext("2d");
    let imgData = inputCtx.getImageData(0, 0, width, height);
    outputCtx.putImageData(
      new ImageData( window.convertPixels(imgData.data, imgData.data.length), width, height)
      , 0, 0);
  }
}
