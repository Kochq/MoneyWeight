import { View, Text, StyleSheet } from "react-native";

interface Transaction {
    id: number;
    title: string;
    amount: number;
    category_id: number;
    subcategory_id: number;
    currency: string;
    payment_method: string;
    exchange_rate: number;
    notes: string;
    date: string; // 2025-01-01 00:00:00
    installment_plan_id: number;
    recurring_payment_id: number;
    payment_number: number;
}

export default function TransactionCard({ t }: { t: Transaction }) {
    const day = t.date.split(" ")[0];

    return (
        <View style={styles.container}>
            <View style={styles.title}>
                <Text style={{ color: "#fff" }}>{t.title}</Text>
            </View>
            <View style={styles.desc}>
                <Text style={{ color: "#0F0" }}>{t.amount}</Text>
                <Text style={{ color: "#fff" }}>{day}</Text>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        borderRadius: 5,
        flexDirection: "row",
        borderWidth: 0.1,
        borderColor: "#fff",
        width: 330,
        paddingVertical: 10,
        paddingHorizontal: 20,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#222",
        gap: 10,
    },

    title: {
        flex: 1,
    },

    desc: {
        flexDirection: "column",
        alignItems: "flex-end",
    },
});
