import JoinForm from "./JoinForm";
import Header from "./Header";
import "./styles.css";
import Conference from "./Conference";
import { useEffect, useState } from "react";
import {
  selectIsConnectedToRoom,
  useHMSActions,
  useHMSStore
} from "@100mslive/hms-video-react";
import Footer from "./Footer";
import "./wasm_exec"

export default function App() {
  const isConnected = useHMSStore(selectIsConnectedToRoom);
  const hmsActions = useHMSActions();
  const [wasm, setWasm] = useState(null);
  const [goInstance, setGoInstance] = useState(new Go());

  useEffect(async () => {
    WebAssembly.instantiateStreaming(fetch("main.wasm"), goInstance.importObject).then(
      async result => {
        await goInstance.run(result.instance);
      })
  },[]);

  

  useEffect(() => {
    window.onunload = () => {
      if (isConnected) {
        hmsActions.leave();
      }
    };
  }, [hmsActions, isConnected]);

  return (
    <div className="App">
      <Header />
      {isConnected ? (
        <>
          <Conference />
          <Footer />
        </>
      ) : (
        <JoinForm />
      )}
    </div>
  );
}
