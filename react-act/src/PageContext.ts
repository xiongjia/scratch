import React, { createContext, useContext } from "react";

export type PageContextType = {
    name: string;
    setName: (name: string) => void;
}

export const PageContext = createContext<PageContextType>({
    name: "test context",
    setName: name => console.log(`set name {name}`)
});

export const usePage = () => useContext(PageContext);