import { useTheme } from "@/theme/ThemeContext";
import Screen from "@/components/Screen";
import {
    View,
    TextInput,
    Text,
    Pressable,
    StyleSheet,
    KeyboardAvoidingView,
    Platform,
    ScrollView,
} from "react-native";
import { useState } from "react";
import { Picker } from "@react-native-picker/picker";
import { CategoryType } from "@/types";

export default function AddCategory() {
    const { colors } = useTheme();
    const [name, setName] = useState<string>("");
    const [icon, setIcon] = useState<string>("");
    const [budgetLimit, setBudgetLimit] = useState("");
    const [selectedTypeCategory, setSelectedTypeCategory] = useState<CategoryType>();
    const apiUrl = process.env.EXPO_PUBLIC_API_BASE;

    const addCategory = async () => {
        if (!name || !budgetLimit || !selectedTypeCategory || !icon) {
            alert("Por favor completa todos los campos");
            return;
        }

        const url = apiUrl + "/api/categories";

        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name: name,
                    budget_limit: Number(budgetLimit),
                    type: selectedTypeCategory,
                    icon: icon
                }),
            });

            const data = await response.json();
            alert(data.status);

            // Limpiar formulario
            setName("");
            setBudgetLimit("");
            setIcon("");
        } catch (e) {
            alert("Error al guardar la transacción");
            console.error(e);
        }
    };

    const styles = StyleSheet.create({
        container: {
            flex: 1,
            backgroundColor: colors.background,
        },

        card: {
            backgroundColor: colors.surface,
            minWidth: "100%",
            borderRadius: 15,
            padding: 20,
            boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
        },

        inputContainer: {
            marginBottom: 16,
        },

        label: {
            fontSize: 14,
            color: colors.textSecondary,
            marginBottom: 8,
            fontWeight: "500",
        },

        input: {
            backgroundColor: colors.background,
            borderRadius: 10,
            padding: 12,
            fontSize: 16,
            color: colors.textPrimary,
            borderWidth: 1,
            borderColor: colors.border,
        },

        picker: {
            backgroundColor: colors.background,
            borderRadius: 10,
            marginBottom: 16,
            borderWidth: 1,
            borderColor: colors.border,
        },

        addButton: {
            backgroundColor: colors.primary,
            borderRadius: 12,
            padding: 16,
            alignItems: "center",
            marginTop: 10,
        },

        addButtonText: {
            color: colors.surface,
            fontSize: 16,
            fontWeight: "600",
        },

        moneyInput: {
            fontSize: 24,
            textAlign: "center",
            color: colors.primary,
            fontWeight: "bold",
        },

        titleInput: {
            textAlign: "center",
            color: colors.textPrimary,
        },
    });
    return (
        <Screen>
            <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                style={styles.container}
            >
                <ScrollView>
                    <View style={styles.card}>
                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Title</Text>
                            <TextInput
                                style={[styles.input, styles.moneyInput]}
                                placeholder="Name of the category..."
                                value={name}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setName}
                            />
                        </View>

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Budget limit</Text>
                            <TextInput
                                style={[styles.input, styles.titleInput]}
                                placeholder="$0.00"
                                value={budgetLimit}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setBudgetLimit}
                            />
                        </View>

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Icon</Text>
                            <TextInput
                                style={[styles.input, styles.titleInput]}
                                placeholder="Icon of the category..."
                                value={icon}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setIcon}
                            />
                        </View>

                        <Text style={styles.label}>Category type</Text>
                        <View style={styles.picker}>
                            <Picker
                                selectedValue={selectedTypeCategory}
                                onValueChange={setSelectedTypeCategory}
                                style={{ color: colors.textPrimary }}
                                prompt="Select a category type"
                            >
                                <Picker.Item
                                    label="Select category type"
                                    value={null}
                                />
                                <Picker.Item
                                    key={CategoryType.income}
                                    label={CategoryType.income}
                                    value={CategoryType.income}
                                />
                                <Picker.Item
                                    key={CategoryType.expense}
                                    label={CategoryType.expense}
                                    value={CategoryType.expense}
                                />
                                <Picker.Item
                                    key={CategoryType.investment}
                                    label={CategoryType.investment}
                                    value={CategoryType.investment}
                                />
                            </Picker>
                        </View>

                        <Pressable
                            style={styles.addButton}
                            onPress={addCategory}
                        >
                            <Text style={styles.addButtonText}>
                                Agregar categoria
                            </Text>
                        </Pressable>
                    </View>
                </ScrollView>
            </KeyboardAvoidingView>
        </Screen>
    );
}
