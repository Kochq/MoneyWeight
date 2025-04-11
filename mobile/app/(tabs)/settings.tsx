import BotonGrandote from "@/components/BotonGrandote";
import { useTheme } from "@/theme/ThemeContext";
import { Link } from "expo-router";
import React from "react";
import { Pressable, StyleSheet, Text, View } from "react-native";
import { Ionicons } from "@expo/vector-icons";

export default function Settings() {
    const { colors } = useTheme();
    return (
        <View style={styles.container}>
            <View style={styles.botones}>
                <BotonGrandote to="/AddMoney" color={colors.primary}>
                    <Ionicons name="person-outline" size={120} color={colors.primary} />
                </BotonGrandote>
                <BotonGrandote to="/AddMoney" color={colors.primary}>
                    <Ionicons name="git-branch-outline" size={120} color={colors.primary} />
                </BotonGrandote>
            </View>
            <View style={styles.botones}>
                <BotonGrandote to="/AddCategory" color={colors.primary}>
                    <Ionicons name="layers-outline" size={120} color={colors.primary} />
                </BotonGrandote>
                <BotonGrandote to="/AddMoney" color={colors.primary}>
                    <Ionicons name="list-outline" size={120} color={colors.primary} />
                </BotonGrandote>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "center",
        backgroundColor: "#121212",
        alignItems: "center",
    },

    botones: {
        gap: 15,
        width: "100%",
        marginBottom: 20,
        flexDirection: "row",
        justifyContent: "space-evenly",
    },

    link: {
        color: "blue",
        backgroundColor: "lightblue",
        padding: 10,
        borderRadius: 5,
        marginBottom: 10,
    },

    title: {
        fontSize: 20,
        fontWeight: "bold",
    },

    text: {
        fontSize: 16,
        marginBottom: 10,
    },
});
