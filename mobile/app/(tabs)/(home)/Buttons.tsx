import { Pressable, StyleSheet } from "react-native";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { Link } from "expo-router";

export const Buttons = () => {
    return (
        <Link asChild href="/AddMoney">
            <Pressable style={styles.buttons}>
                <MaterialIcons name="add" size={50} color="black" />
            </Pressable>
        </Link>
    );
};

const styles = StyleSheet.create({
    buttons: {
        borderWidth: 3,
        borderRadius: 100,
        marginBottom: 20,
        alignItems: "center",
    },
});
