//go:build android

package bluetooth_android

/*
##include <jni>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static inline void copyToError(char* errorMsgv, char** error){
    *error = (char*)malloc((strlen(errorMsg) + 1) * sizeof(char));
    strcpy(*error, errorMsg);
}

jobject getBluetoothOutputStream(uintptr_t env, jobject clientSocket, char** errorMsg) {
    JNIEnv *envPtr = (JNIEnv*)env;
    // Get the BluetoothSocket class
    jclass socketClass = (*envPtr)->GetObjectClass(envPtr, clientSocket);
    if (socketClass == NULL) {
        copyToError( "Failed to get the BluetoothSocket class", errorMsg);
        return NULL;
    }

    // Get the getOutputStream method
    jmethodID getOutputStreamMethod = (*envPtr)->GetMethodID(envPtr, socketClass, "getOutputStream", "()Ljava/io/OutputStream;");
    if (getOutputStreamMethod == NULL) {
        copyToError( "Failed to get getOutputStream method", errorMsg);
        (*envPtr)->DeleteLocalRef(envPtr, socketClass);
        return NULL;
    }

    // Call the getOutputStream method
    jobject outputStream = (*envPtr)->CallObjectMethod(envPtr, clientSocket, getOutputStreamMethod);
    if (outputStream == NULL) {
        copyToError( "Failed to get OutputStream",  errorMsg);
        (*envPtr)->DeleteLocalRef(envPtr, socketClass);
    }

    // Release local references
    (*envPtr)->DeleteLocalRef(envPtr, socketClass);

    return outputStream;
}

jobject getBluetoothInputStream(uintptr_t env, jobject clientSocket, char** errorMsg) {
    JNIEnv *envPtr = (JNIEnv*)env;
    // Get the BluetoothSocket class
    jclass socketClass = (*envPtr)->GetObjectClass(envPtr, clientSocket);
    if (socketClass == NULL) {
        copyToError( "Failed to get the BluetoothSocket class", errorMsg);
        return NULL;
    }

    // Get the getInputStream method
    jmethodID getInputStreamMethod = (*envPtr)->GetMethodID(envPtr, socketClass, "getInputStream", "()Ljava/io/InputStream;");
    if (getInputStreamMethod == NULL) {
        copyToError( "Failed to get getInputStream method", envPtr);
        (*envPtr)->DeleteLocalRef(envPtr, socketClass);
        return NULL;
    }

    // Call the getInputStream method
    jobject inputStream = (*envPtr)->CallObjectMethod(envPtr, clientSocket, getInputStreamMethod);
    if (inputStream == NULL) {
        copyToError( "Failed to get InputStream", errorMsg);
    }

    // Release local references
    (*envPtr)->DeleteLocalRef(envPtr, socketClass);

    return inputStream;
}

void closeStream(JNIEnv* env, jobject stream, char** errorMsg) {
    JNIEnv *envPtr = (JNIEnv*)env;
	//get stream class
    jclass streamClass = (*envPtr)->GetObjectClass(envPtr, stream);
    if (streamClass == NULL) {
        copyToError( "Failed to get the streamClass class", errorMsg);
        return NULL;
    }

    jmethodID closeMethod = (*envPtr)->GetMethodID(envPtr, streamClass, "close", "()V");
    if (closeMethod != NULL) {
        (*envPtr)->CallVoidMethod(envPtr, stream, closeMethod);
    } else {
        copyToError("Failed to close output stream", errorMsg);
    }
    (*env)->DeleteLocalRef(env, streamClass);
    (*env)->DeleteLocalRef(env, stream);
}

char* readFromInputStream(uintptr_t env, jobject inputStream, int size, int* count, char** errorMsg) {
    JNIEnv* envPtr = (JNIEnv*)env;

    // Get the InputStream class
    jclass inputStreamClass = (*envPtr)->GetObjectClass(envPtr, inputStream);
    if (inputStreamClass == NULL) {
        copyToError( "Failed to get the InputStream class", errorMsg);
        return NULL;
    }

    // Get the read method from InputStream
    jmethodID readMethod = (*envPtr)->GetMethodID(envPtr, inputStreamClass, "read", "([B)I");
    if (readMethod == NULL) {
        *count = -1;
        (*envPtr)->DeleteLocalRef(envPtr, inputStreamClass);
        copyToError("Failed to get read method", errorMsg);
        return NULL;
    }

    // Create a byte array for reading data
    jbyteArray byteArray = (*envPtr)->NewByteArray(envPtr, size);
    if (byteArray == NULL) {
        *count = -1;
        (*envPtr)->DeleteLocalRef(envPtr, inputStreamClass);
        copyToError( "Failed to create byte array", errorMsg);
        return NULL;
    }

    // Call the read method
    jint bytesRead = (*envPtr)->CallIntMethod(envPtr, inputStream, readMethod, byteArray);
    if (bytesRead < 0) {
        *count = -1;
        (*envPtr)->DeleteLocalRef(envPtr, inputStreamClass);
        (*envPtr)->DeleteLocalRef(envPtr, byteArray);
        copyToError("Failed to read from input stream", errorMsg);
        return NULL;
    }

    // Copy the bytes from the byte array to a new buffer
    jbyte* byteBuffer = (*envPtr)->GetByteArrayElements(envPtr, byteArray, NULL);
    if (byteArray == NULL) {
        *count = -1;
        (*envPtr)->DeleteLocalRef(envPtr, inputStreamClass);
        (*envPtr)->DeleteLocalRef(envPtr, byteArray);
        copyToError( "Failed to get byte array address", errorMsg);
        return NULL;
    }
    char* buffer = (char*)malloc(bytesRead * sizeof(char));
    memcpy(buffer, byteBuffer, bytesRead);

    // Release memory
    (*envPtr)->DeleteLocalRef(envPtr, inputStreamClass);
    (*envPtr)->DeleteLocalRef(envPtr, byteArray);
    (*envPtr)->ReleaseByteArrayElements(envPtr, byteArray, byteBuffer, 0);

    // Set the count
    *count = bytesRead;

    return buffer;
}

void writeToOutputStream(uintptr_t env, jobject outputStream, const char* buffer, int size, char** errorMsg) {
    JNIEnv* envPtr = (JNIEnv*)env;

    // Get the OutputStream class
    jclass outputStreamClass = (*envPtr)->GetObjectClass(envPtr, outputStream);
    if (outputStreamClass == NULL) {
        copyToError( "Failed to get the OutputStream class", errorMsg);
        return;
    }

    // Get the write method from OutputStream
    jmethodID writeMethod = (*envPtr)->GetMethodID(envPtr, outputStreamClass, "write", "([B)V");
    if (writeMethod == NULL) {
        (*envPtr)->DeleteLocalRef(envPtr, outputStreamClass);
        copyToError("Failed to get write method", errorMsg);
        return;
    }

    // Create a byte array from the buffer
    jbyteArray byteArray = (*envPtr)->NewByteArray(envPtr, size);
    if (byteArray == NULL) {
        (*envPtr)->DeleteLocalRef(envPtr, outputStreamClass);
        copyToError("Failed to create byte array", errorMsg);
        return;
    }
    (*envPtr)->SetByteArrayRegion(envPtr, byteArray, 0, size, (jbyte*)buffer);

    // Call the write method
    (*envPtr)->CallVoidMethod(envPtr, outputStream, writeMethod, byteArray);

    // Release memory
    (*envPtr)->DeleteLocalRef(envPtr, outputStreamClass);
    (*envPtr)->DeleteLocalRef(envPtr, byteArray);
    return;
}
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

type ReadWriterBluetooth struct {
	in, out C.jobject
}

func (r *ReadWriterBluetooth) Read(env uintptr, p []byte) (n int, err error) {
	if p == nil || len(p) == 0 {
		return -1, errors.New("empty buffer")
	}
	var errMsgC *C.char
	var result C.int
	dataC := C.readFromInputStream(C.uintptr_t(env), r.in, C.int(cap(p)), &result, &errMsgC)
	if errMsgC != nil {
		err = errors.New(C.GoString(errMsgC))
		C.free(unsafe.Pointer(errMsgC))
		n = -1
		return
	}
	n = int(result)
	if n < 1 {
		return
	}
	dataGo := C.GoBytes(unsafe.Pointer(dataC), n)
	copy(p, dataGo)
	C.free(unsafe.Pointer(dataC))
	return
}

func (r *ReadWriterBluetooth) Write(env uintptr, p []byte) (err error) {
	if p == nil || len(p) == 0 {
		return errors.New("empty buffer")
	}
	var errMsgC *C.char
	C.writeToOutputStream(C.uintptr_t(env), r.in, (*C.char)(unsafe.Pointer(&p[0])), C.int(cap(p)), &errMsgC)
	runtime.KeepAlive(p)
	if errMsgC != nil {
		err = errors.New(C.GoString(errMsgC))
		C.free(unsafe.Pointer(errMsgC))
	}
	return
}

func (r *ReadWriterBluetooth) Close(env uintptr) error {
	return errors.Join(r.close(env, r.in), r.close(env, r.out))
}

func (r *ReadWriterBluetooth) close(env uintptr, stream C.jobject) error {
	var errMsgC *C.char
	C.closeBluetoothServerSocket(C.uintptr_t(env), stream, &errMsgC)
	if errMsgC != nil {
		err := errors.New(C.GoString(errMsgC))
		C.free(unsafe.Pointer(errMsgC))
		return err
	}
	return nil
}
