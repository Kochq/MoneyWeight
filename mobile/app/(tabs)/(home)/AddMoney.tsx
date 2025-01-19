import {
    View,
    TextInput,
    Text,
    Pressable,
    KeyboardAvoidingView,
    Platform,
    ScrollView,
} from "react-native";
import Screen from "./Screen";
import { useEffect, useState } from "react";
import { Picker } from "@react-native-picker/picker";
import { Category } from "@/types";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../../../theme/ThemeContext";
import { DateTimePickerAndroid } from "@react-native-community/datetimepicker";
import { AddMoneyStyles } from "@/styles/styles";

export default function AddMoney() {
    const { colors } = useTheme();
    const [name, setName] = useState<string>("");
    const [money, setMoney] = useState("");
    const [categories, setCategories] = useState<Category[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<Category>();
    const [date, setDate] = useState<Date>(new Date());
    const styles = AddMoneyStyles(colors);

    const onChange = (event: any, selectedDate: Date) => {
        const currentDate = selectedDate;
        setDate(currentDate);
    };

    const showMode = (currentMode: any) => {
        DateTimePickerAndroid.open({
            value: date,
            onChange,
            mode: currentMode,
            is24Hour: true,
        });
    };

    const showDatepicker = () => showMode("date");
    const showTimepicker = () => showMode("time");

    const addTransaction = async () => {
        if (!name || !money) {
            alert("Por favor completa todos los campos");
            return;
        }

        const formatDate = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()} ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;

        try {
            const response = await fetch(
                "http://serkq.org:8080/api/transactions",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        title: name,
                        amount: Number(money),
                        category_id: selectedCategory || 1,
                        subcategory_id: 1,
                        currency: "ARS",
                        payment_method: "cosa",
                        date: formatDate,
                        exchange_rate: 1500,
                        notes: "esto es una nota",
                        from_account_id: 1,
                    }),
                },
            );

            const data = await response.json();
            alert(data.status);

            // Limpiar formulario
            setName("");
            setMoney("");
            setDate(new Date());
        } catch (e) {
            alert("Error al guardar la transacción");
            console.error(e);
        }
    };

    const getCategories = async () => {
        try {
            const response = await fetch(
                "http://serkq.org:8080/api/categories",
            );
            const data = await response.json();
            setCategories(data.data);
        } catch (e) {
            console.error("Error al obtener categorías:", e);
        }
    };

    useEffect(() => {
        getCategories();
    }, []);

    return (
        <Screen>
            <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                style={styles.container}
            >
                <ScrollView>
                    <View style={styles.card}>
                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Monto</Text>
                            <TextInput
                                style={[styles.input, styles.moneyInput]}
                                placeholder="$0.00"
                                keyboardType="numeric"
                                value={money}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setMoney}
                            />
                        </View>

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Titulo</Text>
                            <TextInput
                                style={[styles.input, styles.titleInput]}
                                placeholder="What did you buy?"
                                value={name}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setName}
                            />
                        </View>

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Categoría</Text>
                            <View style={styles.picker}>
                                <Picker
                                    selectedValue={selectedCategory}
                                    onValueChange={setSelectedCategory}
                                    style={{ color: colors.textPrimary }}
                                >
                                    <Picker.Item
                                        label="Select a category"
                                        value={null}
                                    />
                                    {categories.map((category) => (
                                        <Picker.Item
                                            key={category.id}
                                            label={category.name}
                                            value={category.id}
                                        />
                                    ))}
                                </Picker>
                            </View>
                        </View>

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Fecha y Hora</Text>
                            <View style={styles.dateContainer}>
                                <Text style={styles.dateText}>
                                    {date.toLocaleString()}
                                </Text>
                                <Pressable
                                    style={styles.dateButton}
                                    onPress={showDatepicker}
                                >
                                    <Ionicons
                                        name="calendar"
                                        size={24}
                                        color={colors.primary}
                                    />
                                </Pressable>
                                <Pressable
                                    style={styles.dateButton}
                                    onPress={showTimepicker}
                                >
                                    <Ionicons
                                        name="time"
                                        size={24}
                                        color={colors.primary}
                                    />
                                </Pressable>
                            </View>
                        </View>

                        <Pressable
                            style={styles.addButton}
                            onPress={addTransaction}
                        >
                            <Text style={styles.addButtonText}>
                                Agregar Transacción
                            </Text>
                        </Pressable>
                    </View>
                </ScrollView>
            </KeyboardAvoidingView>
        </Screen>
    );
}
