import { StyleSheet, View } from "react-native";
import PlusImg from "@expo/vector-icons/Entypo";
import MinusImg from "@expo/vector-icons/Ionicons";
import { useTheme } from "../../../theme/ThemeContext";
import { Boton } from "./Boton";

export const Buttons = () => {
    const { colors } = useTheme();

    return (
        <View style={styles.container}>
            <Boton to="/AddMoney" color={colors.moneyIn}>
                <PlusImg name="plus" size={50} color={colors.moneyIn} />
            </Boton>

            <Boton to="/AddMoney" color={colors.moneyOut}>
                <MinusImg name="remove" size={50} color={colors.moneyOut} />
            </Boton>

            <Boton to="/AddMoney" color={colors.installment}>
                <PlusImg name="plus" size={50} color={colors.installment} />
            </Boton>

            <Boton to="/AddMoney" color={colors.recurring}>
                <PlusImg name="plus" size={50} color={colors.recurring} />
            </Boton>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flexDirection: "row",
        justifyContent: "space-between",
        width: "100%",
    },
});
