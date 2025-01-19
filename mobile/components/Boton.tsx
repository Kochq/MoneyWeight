import { Href, Link } from "expo-router";
import { ReactNode } from "react";
import { Pressable } from "react-native";

export default function Boton({
    children,
    to,
    color,
}: {
    children: ReactNode;
    to: Href;
    color: string;
}) {
    return (
        <Link asChild href={to}>
            <Pressable
                style={{
                    borderWidth: 3,
                    borderRadius: 100,
                    marginBottom: 20,
                    alignItems: "center",
                    borderColor: color,
                }}
            >
                {children}
            </Pressable>
        </Link>
    );
}
