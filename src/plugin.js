import { GrayscalePlugin } from "./plugins/grayscalePlugin";
import {
  selectIsLocalVideoPluginPresent,
  useHMSActions,
  useHMSStore,
} from "@100mslive/hms-video-react";
import React from "react";
import { BrightnessPlugin } from "./plugins/brightnessPlugin";

export const grayScalePlugin = new GrayscalePlugin();
export const brightnessPlugin = new BrightnessPlugin();
export function PluginButton({ plugin, name, framerate }) {
  const isPluginAdded = useHMSStore(
    selectIsLocalVideoPluginPresent(plugin.getName())
  );
  const hmsActions = useHMSActions();

  const togglePluginState = async () => {
    if (!isPluginAdded) {
      await hmsActions.addPluginToVideoTrack(plugin, framerate);
    } else {
      await hmsActions.removePluginFromVideoTrack(plugin);
    }
  };

  return (
    <button id="grayscale-btn" className="btn" onClick={togglePluginState}>
      {`${isPluginAdded ? "Remove" : "Add"} ${name}`}
    </button>
  );
}
