import { View, StyleSheet } from "react-native";
import React from "react";
import { useTheme } from "../../../theme/ThemeContext";

interface ScreenProps {
    children: React.ReactNode;
}

export default function Screen({ children }: ScreenProps) {
    const { colors, toggleTheme } = useTheme();

    return (
        <View
            style={[styles.container, { backgroundColor: colors.background }]}
        >
            {children}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: "center",
        paddingTop: 16,
        backgroundColor: "#121212",
        paddingHorizontal: 12,
    },
});
