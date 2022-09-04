import { createContext, useState } from "react";

const AppContext = createContext();

export function AppWrapper({ children }) {
  const [formTypem, formTypeSet] = useState("free");

  const changeFormType = (formTypem) => {
    formTypeSet(formTypem);
  };

  return (
    <AppContext.Provider value={{ formTypem, changeFormType }}>
      {children}
    </AppContext.Provider>
  );
}

export default AppContext;
