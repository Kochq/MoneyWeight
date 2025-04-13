import { StyleSheet, View } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../theme/ThemeContext";
import Boton from "./Boton";

export default function HomeButtons() {
    const { colors } = useTheme();

    return (
        <View style={styles.container}>
            <Boton to="/AddMoney" color={colors.moneyIn}>
                <Ionicons name="add" size={50} color={colors.moneyIn} />
            </Boton>

            <Boton to="/AddMoney" color={colors.moneyOut}>
                <Ionicons name="remove" size={50} color={colors.moneyOut} />
            </Boton>

            <Boton to="/AddMoney" color={colors.installment}>
                <Ionicons name="crop" size={50} color={colors.installment} />
            </Boton>

            <Boton to="/AddMoney" color={colors.recurring}>
                <Ionicons name="locate" size={50} color={colors.recurring} />
            </Boton>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: "row",
        justifyContent: "space-between",
        width: "100%",
    },
});
