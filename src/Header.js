import {
  selectIsConnectedToRoom,
  useHMSActions,
  useHMSStore,
} from "@100mslive/hms-video-react";
import React from "react";
import { brightnessPlugin, grayScalePlugin, PluginButton } from "./plugin";

function Header() {
  const isConnected = useHMSStore(selectIsConnectedToRoom);
  const hmsActions = useHMSActions();
  

  return (
    <header>
      <img
        className="logo"
        src="https://ashwins93.app.100ms.live/static/media/100ms_logo.3cfd8818.svg"
        alt="logo"
      />
      {isConnected && (
        <div>
          <PluginButton plugin={brightnessPlugin} name={"Brightness"} framerate={15}/>
          <PluginButton plugin={grayScalePlugin} name={"Grayscale"} framerate={15}/>
          <button
            id="leave-btn"
            className="btn btn-danger"
            onClick={() => hmsActions.leave()}
          >
            Leave Room
          </button>
        </div>
      )}
    </header>
  );
}

export default Header;
