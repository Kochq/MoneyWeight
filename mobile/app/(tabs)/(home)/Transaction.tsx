import { Transaction } from "@/types";
import { View, Text, StyleSheet, Image } from "react-native";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";

export default function TransactionCard({ t }: { t: Transaction }) {
    const day = t.date.split(" ")[0];

    return (
        <View style={styles.container}>
            <View style={styles.titles}>
                <MaterialIcons name="category" size={24} color="black" />
                <View style={styles.title}>
                    <Text
                        style={{ color: "#fff" }}
                        numberOfLines={1}
                        ellipsizeMode="tail"
                    >
                        {t.title}
                    </Text>
                </View>
            </View>
            <View style={styles.desc}>
                <Text style={{ color: "#0F0" }}>{`$ ${t.amount}`}</Text>
                <Text style={{ color: "#fff" }}>{day}</Text>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    titles: {
        flexDirection: "row",
        alignItems: "center",
        gap: 10,
    },

    container: {
        borderRadius: 5,
        flexDirection: "row",
        borderWidth: 0.1,
        borderColor: "#fff",
        width: 350,
        paddingVertical: 10,
        paddingHorizontal: 20,
        alignItems: "center",
        backgroundColor: "#222",
        justifyContent: "space-between",
    },

    title: {
        width: "60%",
    },

    desc: {
        flexDirection: "column",
        alignItems: "flex-end",
    },

    category: {
        color: "#fff",
        fontSize: 12,
        paddingVertical: 3,
        paddingHorizontal: 10,
        borderWidth: 0.5,
        borderColor: "#fff",
        borderRadius: 5,
    },
});
